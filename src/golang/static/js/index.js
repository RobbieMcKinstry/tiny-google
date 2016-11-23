var commands = {
    search: 'search ',
    help:   'help',
    upload: 'upload'
};

var url = '/tiny-google';

$(function() {

    function search(cmd, term) {
        // Make a GET request to the /tiny-google URL for the particular search term
        $.getJSON(url, { search_query: cmd }, function(data) {

            var spark = data.spark;
            term.echo("Spark Runtime: " + spark.runtime + "\n");
        });
    }

    function help(cmd, term) {
        term.echo("Type 'upload', 'search <phrase>', or 'help' \n");
    }

    function upload(cmd, term) {
        $.post(url, { document: null } );
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
