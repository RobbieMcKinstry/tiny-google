var commands = {
    search: 'search ',
    help:   'help',
    upload: 'upload'
};

var url = '/tiny-google';

$(function() {

    function compareLinks(a, b) {
        return a.frequency - b.frequency;
    } 

    function search(cmd, term) {

        // Make a GET request to the /tiny-google URL for the particular search term
        $.getJSON(url, { search_query: cmd }, function(data) {

            term.echo("System time: " + data.Time);

            var documents = data.Links[cmd].sort(compareLinks);
            documents.forEach(docData => {
                term.echo("[[!;;] "                + docData.DocumentPath + "]");
                term.echo("Document frequency: "   + docData.Frequency);
            });
        });
    }

    function help(cmd, term) {
        term.echo("Type 'upload', 'search <phrase>', or 'help' \n");
    }

    function upload(cmd, term) {

        $('#file').simpleUpload("/upload", {
            success: function(data){
                console.log("Successful upload");
                //upload successful 
            },
            error: function(error){
                //upload failed                                              
                console.log("Upload failed");
            }
        });

        term.echo(cmd);
    }

    function cmd_handler(cmd, term) {

        if (cmd.startsWith(commands.search)) {
            search(
                $.trim(cmd.substring(commands.search.length)),
                term
            );
        } else if (cmd.startsWith(commands.help)) {
            help(
                $.trim(cmd.substring(commands.help.length)),
                term
            );
        } else if (cmd.startsWith(commands.upload)) {
            upload(
                $.trim(cmd.substring(commands.upload.length)),
                term
            );
        } else {
            term.echo("Invalid command. Type 'help' if you want to see the commands.\n")
        }
    }

    var options = {
        greetings: null,
        prompt: 'tiny-google $ ',
        onInit: function(term) {
            help(null, term);
        },
    };

    $('body').terminal(cmd_handler, options);
});
