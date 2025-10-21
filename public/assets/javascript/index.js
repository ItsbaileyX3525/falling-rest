console.log("Javascript loaded!"); //Keeping just to ensure javascript doesn't break
const apiContainer = document.getElementById("api-container");

function autoFormat(endpoint, data, requestType) {
    const divContainer = document.createElement("div");
    const divFlex = document.createElement("div");
    const divTitle = document.createElement("div");
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

    divContainer.id = "get-container"
    divFlex.id = "flex-get";
    divTitle.innerText = requestType;
    aEl.target = "_blank";
    aEl.href = `${endpoint}`;
    aEl.id = "clickable";
    aEl.innerText = ` ${endpoint}`
    spanOpen.id = "line";
    spanClose.id = "line";
    spanOpenP.innerText = "{"
    spanCloseP.innerText = "}"

    spanOpen.appendChild(spanOpenP)
    spanClose.appendChild(spanCloseP);
    code.appendChild(spanOpen);
    for (let e = 0; e < dataLines.length; e++) {
        const span = document.createElement("span");
        const p = document.createElement("p");
        span.id = "line";
        p.innerText = `    "${dataLines[e][0]}": "${dataLines[e][1]}"`;

        span.appendChild(p);
        code.appendChild(span);

    }
    code.appendChild(spanClose);
    divFlex.appendChild(divTitle);
    divFlex.appendChild(aEl);
    divContainer.appendChild(divFlex);
    divContainer.appendChild(code);
    apiContainer.appendChild(divContainer);

}

async function returnAPIResponse(endpoint, requestType = "GET") {
    let fetchOptions = {
        method: requestType
    };
    if (requestType !== "GET") {
        fetchOptions.headers = {
            'Content-Type': 'application/json'
        };
    }
    let response = await fetch(endpoint, fetchOptions)

    if (!response.ok) {
        console.log("Something went wrong");
        return false;
    }

    const data = await response.json()

    autoFormat(endpoint, data, requestType)
    return data;


}

document.addEventListener("DOMContentLoaded", async () => {
    const response = await returnAPIResponse("/api/seasonalFacts")
    const response2 = await returnAPIResponse("https://api.flik.host/joke")
    const response3 = await returnAPIResponse("https://api.flik.host/test_post", "POST")
})