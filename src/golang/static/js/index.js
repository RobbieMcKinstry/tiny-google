$(function() {

    var commands = {
        search: 'search ',
        help:   'help ',
        upload: 'upload '
    };

    function search(cmd, term) {
        term.echo(cmd);
    }

    function help(cmd, term) {
        term.echo("Type 'upload', 'search <phrase>', or 'help' \n");
    }

    function upload(cmd, term) {
        term.echo(cmd);
    }

    function cmd_handler(cmd, term) {

        if (cmd.startsWith(commands.search) {
            search(
                cmd.substring(commands.search.length),
                term
            );
        } else if (cmd.startsWith(commands.help)) {
            help(
                cmd.substring(commands.help.length),
                term
            );
        } else if (cmd.startsWith(commands.upload)) {
            upload(
                cmd.substring(commands.upload.length),
                term
            );
        } else {
            term.echo("Invalid command. Type 'help' if you want to see the commands.\n")
        }
    }

    var options = {
        width: 500,
        height: 230,
        prompt: 'tiny-google $',
        greetings: 
        onInit: function(term) {
            term.echo("Type 'upload', 'search <phrase>', or 'help' \n");
        },
    };

    $('#terminal').terminal(cmd_handler, options);
});
