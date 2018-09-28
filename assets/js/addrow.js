$(document).ready(function () {
    var counter = 1;

    $("#addrow").on("click", function () {
        var newRow = $("<tr>");
        var cols = "";

        cols += `
        <td data-label="Gender: ">
            <div class="form-group">
                <div class="form-check form-check-inline">
                    <input class="form-check-input" type="radio" name="gender${counter}" id="inlineRadio${counter}1" value="1">
                    <label class="form-check-label" for="inlineRadio${counter}1">male</label>
                </div>
                <div class="form-check form-check-inline">
                    <input class="form-check-input" type="radio" name="gender${counter}" id="inlineRadio${counter}2" value="2">
                    <label class="form-check-label" for="inlineRadio${counter}2">female</label>
                </div>
                <div class="form-check form-check-inline">
                    <input class="form-check-input" type="radio" name="gender${counter}" id="inlineRadio${counter}3" value="3">
                    <label class="form-check-label" for="inlineRadio${counter}3">other</label>
            </div>
        </td>
        <td data-label="Name: ">
            <input type="name" name="name${counter}"  class="form-control"/>
        </td>
        <td data-label="E-Mail: ">
            <input type="email" class="form-control" id="exampleFormControlInput1" placeholder="name@example.com" name="mail${counter}">
        </td>
        <td>
            <input type="button" class="ibtnDel btn btn-md btn-danger" value="Delete">
        </td>`;
        newRow.append(cols);
        $("table.order-list").append(newRow);
        counter++;
        $("#guestCount").attr("value", counter + 1);
    });



    $("table.order-list").on("click", ".ibtnDel", function (event) {
        $(this).closest("tr").remove();
        counter -= 1;
    });


});

function calculateRow(row) {
    var price = +row.find('input[name^="price"]').val();
}

function calculateGrandTotal() {
    var grandTotal = 0;
    $("table.order-list").find('input[name^="price"]').each(function () {
        grandTotal += +$(this).val();
    });
    $("#grandtotal").text(grandTotal.toFixed(2));
}