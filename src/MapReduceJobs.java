import org.apache.hadoop.util.Tool;
import org.apache.hadoop.mapred.JobConf;
import org.apache.hadoop.mapred.jobcontrol.Job;
import org.apache.hadoop.mapred.jobcontrol.JobControl;
import org.apache.hadoop.conf.Configured;
import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.util.ToolRunner;
import org.apache.hadoop.streaming;

public class MapReduceJobs extends Configured implements Tool{

    public int run(String[] args) throws Exception {
        JobControl jobc = new JobControl("MapReduceJobs");

        //Job that will run the word count python scripts
        String[] word_count = 
        {
            "-mapper", "word_count_mapper.py",
            "-reducer", "word_count_reducer.py",
            "-input", "../books/*",
            "-output", "./outputj"
        };

        //Create a new configurationf for the MapReduce job that will
        //Get the word count of each book
        JobConf word_count_conf = new StreamJob.createJob(word_count);
        Job word_count_job = new Job(word_count_conf);
        word_count_conf.addJob(word_count_job);

        //Do the same for the term python files
        String[] term_setup = 
        {
            "-mapper"   , "term_mapper.py",
            "-reducer"  , "term_reducer.py",
            "-input"    , "./outputj/part-00000",
            "-output"   , "./outputmr"
        };

        JobConf term_conf = new StreamJob.createJob(term_setup);
        Job term_job = new Job(term_conf);
        term_conf.addJob(term_job);

        //Set a dependency so that term_job depends on the output of word_count
        term_job.addDependingJob(word_count_job);



    }

    public static void main(String[] args) {
        //ToolRunner to handle command line options
        int result = ToolRunner.run(new Configuration(), new MapReduceJobs(), args);
        System.exit(result);
    }
    
}