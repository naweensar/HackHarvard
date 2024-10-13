# AIDERS - Your Real-Time Patient Monitoring System

<div align="center">
  <img alt="AIDERS" height="200px" src="https://github.com/naweensar/HackHarvard/blob/main/public/images/aiders-logo.PNG">
</div>

***Team Members***\
Aymane Omari\
Mingyu Shen\
Naween Sarwani\
Richard

# 🚩 Latest Update
13/10/2024

# 🤔 What is AIDERS
AIDERS is a real-time monitoring program powered by a vision model to detect emergency situations like falling and choking incidents in hospitals. The system aims to enhance patient safety by promptly alerting healthcare professionals when a patient is in distress.

**Related Tech:**

- Front end: Go, Handlebars, JavaScript, CSS HTML
- Back end: Python
- Visual model: YOLOv8
  
# 🚀 How does AIDERS help?

The main process of the AIDERS program:

1. **Human Movement Recognition**: Training a vision model based on datasets we build, it can recognize whether people get choked or fall over efficiently using the YOLOv8 training model.

2. **Emergency Treatment**: Once it detects an emergency situation, it will take a screenshot and send it to the doctor or your family with a short description via email and text message.

3. **Next Step**: After an emergency situation occurs, it will ask how the patient feels. If there is no response within 5 seconds, it will automatically contact the doctor or call 911.

# 🎯 Accuracy Improvement

**Confidence Parameter**:
- For each emergency situation AIDERS detects, it will return a screenshot and a JSON file with the following format to the backend:

<div align="center">
  <img width="461" alt="Screenshot 2024-10-13 at 1 11 45 AM" src="https://github.com/user-attachments/assets/861214ac-6ca2-4c5a-83e8-7700d6db3960">
</div>

- The Confidence parameter measures how much the movements match models in our dataset. In case of frequent inaccurate reports, it will automatically filter out results with confidence less than 0.5.

- It will then send the screenshot to an LLM to ask whether there is a person who fell or got choked. After getting a positive answer, AIDERS will report it to doctors or the family.

## Setup Instructions

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/your-hackathon-project.git
   ```

2. **Navigate to the project directory**

   ```bash
   cd your-hackathon-project
   ```

3. **Install dependencies**

   ```bash
   npm install
   ```

4. **Start the application**

   ```bash
   npm start
   ```

5. **Open your browser and visit**

   [http://localhost:3000](http://localhost:3000)
