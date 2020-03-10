//custom max min header filter
var minMaxFilterEditor = function (cell, onRendered, success, cancel, editorParams) {

  var end;

  var container = document.createElement("span");

  //create and style inputs
  var start = document.createElement("input");
  start.setAttribute("type", "number");
  start.setAttribute("placeholder", "Min");
  start.setAttribute("min", 0);
  start.setAttribute("max", 100);
  start.style.padding = "4px";
  start.style.width = "50%";
  start.style.boxSizing = "border-box";

  start.value = cell.getValue();

  function buildValues() {
    success({
      start: start.value,
      end: end.value,
    });
  }

  function keypress(e) {
    if (e.keyCode == 13) {
      buildValues();
    }

    if (e.keyCode == 27) {
      cancel();
    }
  }

  end = start.cloneNode();
  end.setAttribute("placeholder", "Max");

  start.addEventListener("change", buildValues);
  start.addEventListener("blur", buildValues);
  start.addEventListener("keydown", keypress);

  end.addEventListener("change", buildValues);
  end.addEventListener("blur", buildValues);
  end.addEventListener("keydown", keypress);


  container.appendChild(start);
  container.appendChild(end);

  return container;
}

//custom max min filter function
function minMaxFilterFunction(headerValue, rowValue, rowData, filterParams) {
  //headerValue - the value of the header filter element
  //rowValue - the value of the column in this row
  //rowData - the data for the row being filtered
  //filterParams - params object passed to the headerFilterFuncParams property

  if (rowValue) {
    if (headerValue.start != "") {
      if (headerValue.end != "") {
        return rowValue >= headerValue.start && rowValue <= headerValue.end;
      } else {
        return rowValue >= headerValue.start;
      }
    } else {
      if (headerValue.end != "") {
        return rowValue <= headerValue.end;
      }
    }
  }

  return false; //must return a boolean, true if it passes the filter.
}


var table = new Tabulator("#audit-table", {
  height: "500px",
  layout: "fitColumns",
  columns: [
    {title: "Timestamp", field: "timestamp", width: 120, headerFilter: "input"},
    {title: "Topic", field: "topic", width: 150, headerFilter: "input"},
    {title: "Initiator", field: "initiator", variableHeight: "true", headerFilter: "input"},
    {title: "Target", field: "target", headerFilter: "input"},
    {title: "Action", field: "action", width: 120, headerFilter: "input"},
    {title: "Reason", field: "reason", headerFilter: "input"},
  ],
});
$("#fetch-form").submit(function (event) {
  $.ajax({
    method: "POST",
    url: "/logs",
    data: $(this).serialize(),
    success: function (body) {
      audits = body.filter(e => e.action != 7).map(function (e) {
        e.timestamp = moment.unix(e.timestamp).format("YYYY-MM-DD");
        switch (e.action) {
          case 0:
            e.action = "AddTopic"
            break;
          case 1:
            e.action = "AddPub"
            break;
          case 2:
            e.action = "AddSub"
            break;
          case 3:
            e.action = "RevokePub"
            break;
          case 4:
            e.action = "RevokeSub"
            break;
          case 5:
            e.action = "WrongCountry"
            break;
          case 6:
            e.action = "Banned"
            break;
        }
        return e
      });
      table.setData(audits)
    }
  })
  event.preventDefault();
});

$(document).ready(function () {
  $("#date-start").val(moment().subtract(1, 'days').format("YYYY-MM-DD"));
  $("#date-end").val(moment().add(1, 'days').format("YYYY-MM-DD"));
})
