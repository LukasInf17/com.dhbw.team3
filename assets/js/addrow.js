
    $(document).ready(function () {
    var addRowCounter = $("#guestcount").attr("value");

        $("#addrow").on("click", function () {
            var newRow = $("<tr>");
            var cols = "";

            cols += `
        <td data-label="Gender: ">
            <div class="form-group">
                <div class="form-check form-check-inline">
                    <input class="form-check-input" type="radio" name="Gender${addRowCounter}" id="inlineRadio${addRowCounter}1" value="1" form="invitation-form">
                    <label class="form-check-label" for="inlineRadio${addRowCounter}1">male</label>
                </div>
                <div class="form-check form-check-inline">
                    <input class="form-check-input" type="radio" name="Gender${addRowCounter}" id="inlineRadio${addRowCounter}2" value="2" form="invitation-form">
                    <label class="form-check-label" for="inlineRadio${addRowCounter}2">female</label>
                </div>
                <div class="form-check form-check-inline">
                    <input class="form-check-input" type="radio" name="Gender${addRowCounter}" id="inlineRadio${addRowCounter}3" value="3" form="invitation-form" checked="checked" />
                    <label class="form-check-label" for="inlineRadio${addRowCounter}3">diverse</label>
            </div>
        </td>
        <td data-label="Name: ">
            <input type="name" name="Name${addRowCounter}" class="form-control" form="invitation-form"/>
        </td>
        <td data-label="E-Mail: ">
            <input type="email" class="form-control" placeholder="name@example.com" name="Mail${addRowCounter}" form="invitation-form">
        </td>
        <td>
            <input type="button" class="ibtnDel btn btn-md btn-danger" value="Delete">
        </td>`;
            newRow.append(cols);
            $("table.order-list").append(newRow);
            addRowCounter++;
            $("#guestcount").attr("value", addRowCounter);
        });



        $("table.order-list").on("click", ".ibtnDel", function (event) {
            $(this).closest("tr").remove();
            addRowCounter--;
            $("#guestcount").attr("value", addRowCounter);
        });


    });