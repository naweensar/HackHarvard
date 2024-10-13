# Aider - Your Real-Time Patient Monitoring System

<div align="center">
 <img alt="Aider" height="200px" src="https://github.com/naweensar/HackHarvard/blob/main/public/images/aiders-logo.PNG">
</div>

***Team Members***\
Aymane Omari\
Mingyu Shen\
Naween Sarwani\
Richard

# 🚩latest update
13/10/2024

# 🤔What is Adier
Aider is a real-time monitoring program powered by vision model to detect emergency situations like falling and choking incidents in hospitals. The system aims to enhance patient safety by promptly alerting healthcare professionals when a patient is in distress.

**Related Tech:**

- Front end: Go, Handlebars, JavaScript, CSS HTML
- Back end: Python
- Visual model: Yolov8
  
# 🚀How does Aider help?

The main process of Aider program are:

1. **human movement recognition**: training a vision model based on datasets we build, it could regonize whether people get chocked or fall over ground efficiently by using YOLOV8 training model

2. **emergency treatment**: Once it detects there's an emergency situation, it will take the screenshot and send it to the doctor or your family with a short descrption via email and text message. 

3. **Next step**: After an emergency situation occurs, it will ask how does patient feel, after 5 seconds without detecting answers, it will contact his doctor or call 911 automitically


# 🎯Accuracy improvement

**Confidence parameter**:
- for each emergency situation Aider detects, it will return the screentshot and a jason file with following foramt to backend:
<div align="center">
<img width="461" alt="Screenshot 2024-10-13 at 1 11 45 AM" src="https://github.com/user-attachments/assets/861214ac-6ca2-4c5a-83e8-7700d6db3960">
</div>

- The Confidence parameter measures how much does the movements match models in our dataset. In case of frequent inaccurate reports, it will automitically filter out results with confidence less then 0.5
  
- Seondly, it will send the screentshot to LLM and ask whether there is a person who fall over or get choked. After getting a positive answer, Aider will report it to doctors or the family
   
**Future Plan**:
- according to limited time and dataset we have so far, Aider merely supports singe, simple movements that are easy to be detected such as falling and choking. However, with more dataset and more time to train the model, it will be more accurate and able to recognize complicate movements like serizures, change of skin and eye colors which could be a symptom of some cancers. 

## Setup Instructions

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/your-hackathon-project.git

2. **Navigate to the project directory**
cd your-hackathon-project

3. **Install dependencies**
npm install

4. **Start the application**
npm start

5. **Open your browser and visit**
http://localhost:3000

