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
        console.log("Something went wrong fetching email");
        return "Not logged in"
    }

    const data = await response.json()

    console.log(data)
    if (data.success || data.success === "true") {
        return data.email
    } else {
        return "Not logged in"
    }
}

document.addEventListener("DOMContentLoaded", async () => {
    username.innerText = await fetchEmail()
    const APIKey = window.localStorage.getItem("APIKey")
    APIKeyEl.innerText = APIKey
    await returnAPIResponse(`/api/decode?input=dGVzdA==?type=base64?apiKey=${APIKey}`);
})