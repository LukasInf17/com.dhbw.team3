$(document).ready(function () {
    var addRowCounter = parseInt($("#guestcount").attr("value"));

    $("#addrow").on("click", function () {
        var newRow = $("<tr>");
        var cols = "";

        cols += `
        <td class="col-sm-4 gender-radio-buttons" data-label="Gender: ">
            <div class="form-group">
                <div class="form-check form-check-inline">
                    <input class="form-check-input" type="radio" name="gender${addRowCounter}" id="inlineRadio${addRowCounter}1" value="1" form="invitation-form">
                    <label class="form-check-label" for="inlineRadio${addRowCounter}1">male</label>
                </div>
                <div class="form-check form-check-inline">
                    <input class="form-check-input" type="radio" name="gender${addRowCounter}" id="inlineRadio${addRowCounter}2" value="2" form="invitation-form">
                    <label class="form-check-label" for="inlineRadio${addRowCounter}2">female</label>
                </div>
                <div class="form-check form-check-inline">
                    <input class="form-check-input" type="radio" name="gender${addRowCounter}" id="inlineRadio${addRowCounter}3" value="3" form="invitation-form" checked="checked" />
                    <label class="form-check-label" for="inlineRadio${addRowCounter}3">diverse</label>
            </div>
        </td>
        <td class="col-sm-3" data-label="Name: ">
            <input type="name" name="name${addRowCounter}" class="form-control" form="invitation-form"/>
        </td>
        <td class="col-sm-4" data-label="E-Mail: ">
            <input type="email" class="form-control" placeholder="name@example.com" name="mail${addRowCounter}" form="invitation-form">
        </td>
        <td class="col-sm-2">
            <input type="button" class="btn btn-md btn-danger" value="Delete" id="delete${addRowCounter}" />
        </td>`;
        $(`#delete${addRowCounter}`).on("click", (event) => {
            console.log(this.id + "pressed");
            $(this).parent().parent().remove();
            addRowCounter--;
            $("#guestcount").attr("value", addRowCounter);
        });
        newRow.append(cols);
        $("table.order-list").append(newRow);
        addRowCounter++;
        $("#guestcount").attr("value", addRowCounter);
    });
});