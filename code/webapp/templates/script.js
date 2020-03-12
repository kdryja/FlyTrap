let auditTable = new Tabulator("#audit-table", {
  layout: "fitColumns",
  pagination: "local",
  paginationSize: 7,
  columns: [
    {title: "Timestamp", field: "timestamp", width: 160, headerFilter: "input"},
    {title: "Topic", field: "topic", width: 150, headerFilter: "input"},
    {title: "Initiator", field: "initiator", variableHeight: "true", headerFilter: "input"},
    {title: "Target", field: "target", headerFilter: "input"},
    {title: "Action", field: "action", width: 120, headerFilter: "input"},
    {title: "Reason", field: "reason", headerFilter: "input"},
  ],
});
let summaryTable = new Tabulator("#summary-table", {
  layout: "fitColumns",
  pagination: "local",
  paginationSize: 10,
  columns: [
    {title: "Timestamp", field: "timestamp", width: 160, headerFilter: "input"},
    {title: "Topic", field: "topic", width: 150, headerFilter: "input"},
    {title: "Publishers", field: "publishers", formatter: "textarea", headerFilter: "input"},
    {title: "Subscribers", field: "subscribers", formatter: "textarea", headerFilter: "input"},
  ],
});

$('#person-input').on('input', function () {
  $('#topic-input').not(this).prop('disabled', this.value.length)
});
$('#topic-input').on('input', function () {
  $('#person-input').not(this).prop('disabled', this.value.length)
});

$(".query-form").submit(function (event) {
  $.ajax({
    method: "POST",
    url: "/detailed",
    data: $(this).serialize(),
    success: function (body) {
      $("#result-entity").text(body.req);
      $("#result-date").text($("#result-date-input").val());
      $(".results").show();
      
      let pubs = $("#publishing-list");
      let subs = $("#subscribing-list");
      pubs.empty();
      subs.empty();
      body.pubs.forEach(e => {
        $("<li/>").text(e).appendTo(pubs)
      });
      body.subs.forEach(e => {
        $("<li/>").text(e).appendTo(subs)
      });

      var notification = document.querySelector('.mdl-js-snackbar');
      notification.MaterialSnackbar.showSnackbar(
        {
          message: 'Fetch successful'
        }
      );
    },
    error: function (xhr, status, error) {
      var notification = document.querySelector('.mdl-js-snackbar');
      notification.MaterialSnackbar.showSnackbar(
        {
          message: 'Eror - ' + xhr.status + ': ' + xhr.statusText
        }
      );
    },
  });
  event.preventDefault();
});

$(".fetch-form").submit(function (event) {
  $.ajax({
    method: "POST",
    url: "/logs",
    data: $(this).serialize(),
    error: function (xhr, status, error) {
      var notification = document.querySelector('.mdl-js-snackbar');
      notification.MaterialSnackbar.showSnackbar(
        {
          message: 'Eror - ' + xhr.status + ': ' + xhr.statusText
        }
      );
    },
    success: function (body) {
      audits = body.filter(e => e.action != 7).map(function (e) {
        e.timestamp = moment.unix(e.timestamp).format("YYYY-MM-DD HH:mm:ss");
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
      auditTable.setData(audits)

      let summaries = [];
      body.filter(e => e.action == 7).forEach(function (t) {
        let topics = {};
        let time = moment.unix(t.timestamp).format("YYYY-MM-DD HH:mm:ss");
        let pubs = JSON.parse(JSON.parse(t.reason).pubs);
        let subs = JSON.parse(JSON.parse(t.reason).subs);
        Object.keys(pubs).forEach(e => {
          topics[e] = {...topics[e], pubs: pubs[e]};
        });
        Object.keys(subs).forEach(e => {
          topics[e] = {...topics[e], subs: subs[e]};
        });
        Object.keys(topics).forEach(e => {
          summaries.push({
            timestamp: time,
            topic: e,
            publishers: topics[e].pubs ? topics[e].pubs.join('\n') : "",
            subscribers: topics[e].subs ? topics[e].subs.join('\n') : ""
          });
        });
      });
      summaryTable.setData(summaries);

      var notification = document.querySelector('.mdl-js-snackbar');
      notification.MaterialSnackbar.showSnackbar(
        {
          message: 'Fetch successful!'
        }
      );
    }
  })
  event.preventDefault();
});

$(document).ready(function () {
  $(".date-start").val(moment().subtract(1, 'days').format("YYYY-MM-DD"));
  $(".date-today").val(moment().format("YYYY-MM-DD"));
  $(".date-end").val(moment().add(1, 'days').format("YYYY-MM-DD"));
})
