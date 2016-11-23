package spark.jobserver

import com.typesafe.config.{Config, ConfigFactory}
import org.apache.spark._


object SimpleApp extends SparkJob {
    def main(args: Array[String]) {
        val conf = new SparkConf().setMaster("local[4]").setAppName("Simple App")
        val sc = new SparkContext(conf)
        val config = ConfigFactory.parseString("")
        val results = runJob(sc, config)
        println("Result is " + results)
    }
    def validate(sc: SparkContext, config: Config): SparkJobValidation = SparkJobValid

    def runJob(sc: SparkContext, config: Config): Any = {
        val logFile = "/mounted/docker-compose.yaml" // Should be some file on your system
        val logData = sc.textFile(logFile, 2).cache()
        val numAs = logData.filter(line => line.contains("a")).count()
        val numBs = logData.filter(line => line.contains("b")).count()
        println(s"Lines with a: $numAs, Lines with b: $numBs")
        sc.stop()
        s"Count of the letter a: $numAs,\n Count of the letter b: $numBs"
    }
}
