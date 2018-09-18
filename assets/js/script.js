var hammertime = new Hammer(document.getElementsByTagName("body")[0]);
hammertime.on("swiperight", function(ev) {
    $('#sidebar').addClass('active');
});
hammertime.on("swipeleft", function(ev) {
    $('#sidebar').removeClass('active');
});