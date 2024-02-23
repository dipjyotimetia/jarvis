## Setup: _Obtaining and Configuring Your Gemini API Key_

To interact with Google's Gemini models, you'll need an API key. Follow these steps to obtain and set it up as an environment variable:

**1. _Generate Your API Key_**

* Visit [https://ai.google.dev/](https://ai.google.dev/)
* In Google AI Studio, click on the "Get API key" button.
* Choose "Create API key".
* Your unique API key will be generated. Copy and save it securely.

**2. _Set the API Key as an Environment Variable_**

The recommended way to manage your API key is to set it as an environment variable. This avoids hardcoding it within your code. Here's how:

**Linux/macOS:**

```bash
export API_KEY="your_actual_api_key" 
```

**Windows:**

```powershell
$Env:API_KEY = "your_actual_api_key"
```

**Important:** Replace `"your_actual_api_key"` with the key you copied in step 1.

**Verification (Optional)**

To confirm your API key is set correctly, use the following commands:

* **Linux/macOS:** `echo $API_KEY`
* **Windows:** `echo %API_KEY%`

