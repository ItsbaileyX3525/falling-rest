const form = document.getElementById("form");

form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const formData = new FormData(form);
    
    const response = await fetch("/auth/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        credentials: "include",
        body: JSON.stringify({ 
            "email": formData.get("email"),
            "password": formData.get("pwd")
        })
    });

    if (!response.ok) {
        console.log("Login failed");
        return;
    }

    const data = await response.json();

    if (data.success || data.success === 'true') {
        // Store API key from login response
        window.localStorage.setItem("APIKey", data.api_key);
        // Redirect to account page
        window.location.href = "/account";
    } else {
        console.log(data.message);
        alert(data.message || "Login failed");
    }
});