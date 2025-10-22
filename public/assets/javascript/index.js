console.log("Javascript loaded!"); //Keeping just to ensure javascript doesn't break
const apiContainer = document.getElementById("api-container");
const bonusApiContainer = document.getElementById("bonus-api-container");

function createLoadingContainer(endpoint, requestType = "GET", isBonus = false) {
    const divContainer = document.createElement("div");
    divContainer.id = "get-container";
    divContainer.innerText = "Loading...";
    if (!isBonus) {
        apiContainer.appendChild(divContainer);
    } else {
        bonusApiContainer.appendChild(divContainer);
    }
    return divContainer;
}

function autoFormat(endpoint, data, requestType, isBonus = false, existingContainer = null) {
    const divContainer = document.createElement("div");
    const divFlex = document.createElement("div");
    const divTitle = document.createElement("div");
    const getContainer = document.createElement("div");
    const aEl = document.createElement("a");
    const code = document.createElement("code");
    const spanOpen = document.createElement("span");
    const spanOpenP = document.createElement("p");
    const spanClose = document.createElement("span");
    const spanCloseP = document.createElement("p");
    const dataLines = []

    for (const [key, value] of Object.entries(data)) {
        const data = [key, value]
        dataLines.push(data);
    }

    let containerEl = existingContainer;
    if (!containerEl) {
        containerEl = divContainer;
        containerEl.id = "get-container";
        containerEl.innerText = "Loading...";
        if (!isBonus) {
            apiContainer.appendChild(containerEl);
        } else {
            bonusApiContainer.appendChild(containerEl);
        }
    }
    divFlex.id = "flex-get";
    divTitle.innerText = requestType;
    aEl.target = "_blank";
    aEl.href = `${endpoint}`;
    aEl.id = "clickable";
    aEl.innerText = ` ${endpoint}`
    spanOpen.id = "line";
    spanClose.id = "line";
    spanOpen.style.color = "#7C7F93"
    spanClose.style.color = "#7C7F93"
    spanOpenP.innerText = "{"
    spanCloseP.innerText = "}"
    getContainer.id = "request-container"

    spanOpen.appendChild(spanOpenP)
    spanClose.appendChild(spanCloseP);
    code.appendChild(spanOpen);
    for (let e = 0; e < dataLines.length; e++) {
        const span = document.createElement("span");
        const spanKey = document.createElement("span");
        const spanValue = document.createElement("span");
        span.id = "line";
        spanKey.style.color = "#1E66F5";
        spanValue.style.color = "#40A02B"

        spanKey.innerText = `    "${dataLines[e][0]}":`;
        spanValue.innerText = ` "${dataLines[e][1]}"`

        span.appendChild(spanKey);
        span.appendChild(spanValue);
        code.appendChild(span);

    }
    containerEl.innerText = "";
    code.appendChild(spanClose);
    divFlex.appendChild(divTitle);
    divFlex.appendChild(aEl);
    containerEl.appendChild(divFlex);
    getContainer.appendChild(code)
    containerEl.appendChild(getContainer);
}

async function returnAPIResponse(endpoint, requestType = "GET", isBonus = false) {
    let fetchOptions = {
        method: requestType
    };
    if (requestType !== "GET") {
        fetchOptions.headers = {
            'Content-Type': 'application/json'
        };
    }
    const loadingEl = createLoadingContainer(endpoint, requestType, isBonus);

    let response;
    try {
        response = await fetch(endpoint, fetchOptions)
    } catch (err) {
        console.log("Network error:", err);
        loadingEl.innerText = "Network error";
        return false;
    }

    if (!response.ok) {
        console.log("Something went wrong");
        return false;
    }

    const data = await response.json()

    autoFormat(endpoint, data, requestType, isBonus, loadingEl)
    return data;


}

document.addEventListener("DOMContentLoaded", async () => {
    await returnAPIResponse("/api/seasonalFacts")
    await returnAPIResponse("/api/scientificFacts")
    await returnAPIResponse("/api/leavesImages")

    
    //Other endpoints - not related to site
    await returnAPIResponse("https://api.flik.host/joke", "GET", true)
    await returnAPIResponse("https://api.deer.rest/fact", "GET", true)
    await returnAPIResponse("https://api.flik.host/test_post", "POST", true)
})