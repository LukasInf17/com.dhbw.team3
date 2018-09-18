require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
$(() => {
    var hammertime = new Hammer(document.getElementsByTagName("body")[0]);
    hammertime.on("swiperight", function (ev) {
        $('#sidebar').addClass('active');
    });
    hammertime.on("swipeleft", function (ev) {
        $('#sidebar').removeClass('active');
    });
    $('#sidebarCollapse').on('click', function () {
        $('#sidebar').toggleClass('active');
    });
});
