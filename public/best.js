$(document).ready(function () {
  function createStatusBlock(check) {
    const statusBlock = $("<div class='status-block'></div>");

    statusBlock.addClass(
      check.Statuses[0].Status === "OK" ? "status-ok" : "status-fail"
    );

    statusBlock.data("check", check);

    const tooltipContent =
      "<strong>Service:</strong> " +
      check.Title +
      "<br><strong>Description:</strong> " +
      check.Description +
      "<br><strong>Status:</strong> " +
      check.Statuses[0].Status +
      "<br><strong>Latency:</strong> " +
      check.Statuses[0].Latency +
      "<br><strong>Date:</strong> " +
      check.Statuses[0].Date;

    statusBlock.mouseenter(function () {
      const tooltip = $("#tooltip");

      tooltip.html(tooltipContent);

      tooltip.css({
        top: statusBlock.offset().top + 40,
        left: statusBlock.offset().left,
      });

      tooltip.fadeIn();
    });

    statusBlock.mouseleave(function () {
      $("#tooltip").fadeOut();
    });

    return statusBlock;
  }

  function populateGrid(data) {
    const container = $("#status-container");

    container.empty();

    data.forEach(function (check) {
      const statusBlock = createStatusBlock(check);
      container.append(statusBlock);
    });
  }

  function displayFailedStatuses(data) {
    const failedStatuses = [];
    data.forEach(function (check) {
      check.Statuses.forEach(function (status) {
        if (status.Status === "FAIL") {
          failedStatuses.push({
            Service: check.Title,
            Description: check.Description,
            Latency: status.Latency,
            Date: status.Date,
          });
        }
      });
    });

    const failedStatusList = $("#failed-status-list");
    failedStatusList.empty();

    if (failedStatuses.length === 0) {
      $("#show-downtime").text("No offline service(s) history.");
      $("#failed-status-list").hide();
      $("#downtime-logs-title").hide();
    } else {
      failedStatuses.forEach(function (statusInfo) {
        const listItem = $("<li></li>");
        listItem.html(
          `<strong>Service:</strong> ${statusInfo.Service}<br>` +
            `<strong>Status:</strong> Fail<br>` +
            `<strong>Date:</strong> ${statusInfo.Date}`
        );
        failedStatusList.append(listItem);
      });
    }
  }

  $("#failed-status-list").hide();
  $("#downtime-logs-title").hide();
  $("#show-downtime").hide();

  $.get("/api/check", function (data) {
    $("#show-downtime")
      .show()
      .click(function () {
        $("#failed-status-list").toggle();
      });

    populateGrid(data);
    displayFailedStatuses(data);

    $("#preloader").hide();
  });
});
