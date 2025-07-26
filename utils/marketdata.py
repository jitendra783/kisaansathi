import requests
import xmltodict
import pandas as pd

# Step 1: Fetch the XML from API
url = "https://api.data.gov.in/resource/35985678-0d79-46b4-9ed6-6f13308a1d24"
params = {
    "api-key": "579b464db66ec23bdd000001cdd3946e44ce4aad7209ff7b23ac571b",
    "format": "xml"
}

print("📡 Fetching data...")
response = requests.get(url, params=params)
response.raise_for_status()

# Step 2: Parse XML response
data_dict = xmltodict.parse(response.text)

# Step 3: Access correct path
records = data_dict["result"]["records"]["item"]

# Step 4: Convert to DataFrame
df = pd.DataFrame(records)

# Step 5: Export to Excel
output_file = "mandi_prices_data_gov_in.xlsx"
df.to_excel(output_file, index=False)

print(f"✅ Excel saved: {output_file}")
