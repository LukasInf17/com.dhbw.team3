$(document).ready(function () {
    var counter = 0;

    $("#addrow").on("click", function () {
        var newRow = $("<tr>");
        var cols = "";

        cols += '<td><div class="form-group"><div class="form-check form-check-inline"><label class="form-check-label"><input class="form-check-input" type="radio" name="inlineRadioOptions' + counter + '" id="inlineRadio1" value="option1"> male </label></div><div class="form-check form-check-inline"><label class="form-check-label"><input class="form-check-input" type="radio" name="inlineRadioOptions' + counter + '" id="inlineRadio2" value="option2"> female </label></div><div class="form-check form-check-inline"><label class="form-check-label"><input class="form-check-input" type="radio" name="inlineRadioOptions' + counter + '" id="inlineRadio2" value="option3"> none </label></div></td>';
        cols += '<td><input type="name" name="name' + counter + '"  class="form-control"/></td>';
        cols += '<td><input type="email" class="form-control" id="exampleFormControlInput1" placeholder="name@example.com" name="name' + counter + '"></td>';

        cols += '<td><input type="button" class="ibtnDel btn btn-md btn-danger "  value="Delete"></td>';
        newRow.append(cols);
        $("table.order-list").append(newRow);
        counter++;
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
