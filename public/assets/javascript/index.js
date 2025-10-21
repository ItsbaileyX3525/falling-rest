console.log("Javascript loaded!"); //Keeping just to ensure javascript doesn't break
const apiContainer = document.getElementById("api-container");

function autoFormat(endpoint, data) {
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
        console.log(key, value);
        const data = [key, value]
        dataLines.push(data);
    }

    divContainer.id = "get-container"
    divFlex.id = "flex-get";
    divTitle.innerText = "GET";
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
        console.log("Something happening")
        const span = document.createElement("span");
        const p = document.createElement("p");
        span.id = "line";
        console.log(dataLines, dataLines[e][0], dataLines[e][1]);
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

async function returnAPIResponse(endpoint) {
    let response = await fetch(endpoint)

    if (!response.ok) {
        console.log("Something went wrong");
        return false;
    }

    const data = await response.json()

    console.log(data)
    autoFormat(endpoint, data)
    return data;


}

document.addEventListener("DOMContentLoaded", async () => {
    const response = await returnAPIResponse("/api/seasonalFacts")
    const response2 = await returnAPIResponse("https://api.flik.host/joke", true)
})