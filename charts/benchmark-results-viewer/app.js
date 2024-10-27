let currentActiveCores;
let filterFunctionsDebounceTimer;

function renderSidebar() {
  const functionList = document.getElementById("function-list");

  Object.keys(benchmarkResults).forEach((benchmarkName) => {
    const li = document.createElement("li");
    li.textContent = benchmarkName;
    li.addEventListener("click", () => showBenchmark(benchmarkName));
    functionList.appendChild(li);
  });

  document
    .getElementById("search-bar")
    .addEventListener("input", filterFunctions);
}

function filterFunctions() {
  clearTimeout(filterFunctionsDebounceTimer);
  filterFunctionsDebounceTimer = setTimeout(() => {
    const searchTerm = document
      .getElementById("search-bar")
      .value.toLowerCase();
    const functionItems = document.querySelectorAll("#function-list li");

    functionItems.forEach((item) => {
      const funcName = item.textContent.toLowerCase();
      item.style.display = funcName.includes(searchTerm) ? "block" : "none";
    });
  }, 300);
}

function showBenchmark(benchmarkName) {
  const benchmarksContainer = document.getElementById("benchmarks");
  benchmarksContainer.innerHTML = "";

  // Update active state in sidebar
  document.querySelectorAll("#function-list li").forEach((li) => {
    li.classList.toggle("active", li.textContent === benchmarkName);
  });

  const benchmarkElement = document
    .getElementById("benchmark-container-template")
    .content.cloneNode(true);
  const data = benchmarkResults[benchmarkName];
  if (!data) {
    console.error(`No data found for benchmark: ${benchmarkName}`);
    return;
  }

  benchmarkElement.querySelector(".benchmark-title").textContent =
    benchmarkName;

  const scenarioTemplate = document.querySelector(
    ".benchmark-scenario-template"
  );
  const scenarioContainer = benchmarkElement.querySelector(".tab-content");

  data.scenarioResults.forEach((scenario) => {
    const scenarioElement = scenarioTemplate.content.cloneNode(true);
    scenarioElement.querySelector(
      ".benchmark-scenario-description"
    ).textContent = scenario.description;
    const chartElement = scenarioElement.querySelector(
      `.benchmark-scenario-content[data-cores="${currentActiveCores}"] .benchmark-scenario-chart`
    );
    if (chartElement) {
      const scenarioData = {
        benchmarkName,
        sizes: data.sizes,
        durations: scenario.durations,
      };
      renderChart(scenarioData, chartElement, currentActiveCores);
    }
    scenarioContainer.appendChild(scenarioElement);
  });

  updateActiveTab(benchmarkElement);

  benchmarksContainer.appendChild(benchmarkElement);

  window.dispatchEvent(new Event("resize"));
}

function updateActiveTab(element) {
  document.querySelectorAll(".tab-button").forEach((btn) => {
    btn.classList.toggle(
      "active",
      btn.getAttribute("data-cores") === currentActiveCores
    );
  });
  element.querySelectorAll(".benchmark-scenario-content").forEach((content) => {
    content.style.display =
      content.getAttribute("data-cores") === currentActiveCores
        ? "flex"
        : "none";
  });
}

function initTabButtons() {
  document.querySelectorAll(".tab-button").forEach((button) => {
    button.addEventListener("click", () => {
      currentActiveCores = button.getAttribute("data-cores");
      updateActiveTab(document);
      const activeBenchmark = document.querySelector(
        "#function-list li.active"
      );
      if (activeBenchmark) {
        showBenchmark(activeBenchmark.textContent);
      }
    });
  });
}

function handleScroll() {
  const headers = document.querySelectorAll('.benchmark-header');
  headers.forEach(header => {
    const rect = header.getBoundingClientRect();
    if (rect.top <= 0) {
      header.classList.add('scrolled');
    } else {
      header.classList.remove('scrolled');
    }
  });
}

function init() {
  currentActiveCores = document
    .querySelector(".tab-button.active")
    .getAttribute("data-cores");

  renderSidebar();
  initTabButtons();

  showBenchmark(Object.keys(benchmarkResults)[0]);

  // Add scroll event listener
  window.addEventListener('scroll', handleScroll);
}

init();
