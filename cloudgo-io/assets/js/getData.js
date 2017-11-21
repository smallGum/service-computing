$(document).ready(function() {
    $.ajax({
        url: "/"
    }).then(function(data) {
        $('.greeting-author').append(data.author).hide();
        $('.greeting-os').append(data.os).hide();
        $('.greeting-date').append(data.date).hide();
    });

    $.ajax({
        url: "/register"
    }).then(function(data) {
        $('.greeting-id').append(data.id);
        $('.greeting-username').append(data.username);
        $('.greeting-subject').append(data.subject);
        $('.greeting-score').append(data.score);
    });

    $('.info').click(function() {
        $('.info').hide();
        $('.greeting-author').show();
        $('.greeting-os').show();
        $('.greeting-date').show();
    });
});