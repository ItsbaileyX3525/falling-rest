const username = document.getElementById("username");
import { returnAPIResponse } from "./api.js";
const APIKeyEl = document.getElementById("api-key")

async function fetchEmail() {
    const response = await fetch("/auth/me", {
        method: "GET",
        credentials: "include",
        headers: {
            "Accept" : "application/json"
        }
    })

    if (!response.ok) {
        localStorage.clear()
        console.log("Something went wrong fetching email");
        return "Not logged in"
    }

    const data = await response.json()

    console.log(data)
    if (data.success || data.success === "true") {
        return data.email
    } else {
        localStorage.clear()
        return "Not logged in"
    }
}

document.addEventListener("DOMContentLoaded", async () => {
    username.innerText = await fetchEmail()
    const APIKey = window.localStorage.getItem("APIKey") || "Not logged in"
    APIKeyEl.innerText = APIKey
    await returnAPIResponse(`/api/decode?input=YmFzZTY0IGVuY29kZWQgdGV4dA==?type=base64?apiKey=${APIKey}`);
    await returnAPIResponse(`/api/fallPeople?apiKey=${APIKey}`);
    await returnAPIResponse(`/api/fallQuotes?apiKey=${APIKey}`);
})