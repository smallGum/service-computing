$(document).ready(function() {
    $.ajax({
        url: "/"
    }).then(function(data) {
        $('.greeting-author').append(data.author).hide();
        $('.greeting-os').append(data.os).hide();
        $('.greeting-date').append(data.date).hide();
    });

    $('.info').click(function() {
        $('.info').hide();
        $('.greeting-author').show();
        $('.greeting-os').show();
        $('.greeting-date').show();
    });
});