import requests
import cv2
from face_recognition import FaceRecognition
import base64
from face_recognition import FaceRecognition
from sklearn.model_selection import train_test_split
from sklearn.metrics import classification_report, roc_curve, precision_recall_curve, roc_auc_score, accuracy_score
import matplotlib.pyplot as plt
import os
import glob
import pandas as pd
import random
import numpy as np
import sys
import cv2
import base64
from tqdm import tqdm
import requests
from pprint import pprint
import shutil
import time
from parinya import LINE


MODEL_PATH = "model_v1.pkl"
# Initialize the FaceRecognition class
fr = FaceRecognition()

# Load the saved model
fr.load(MODEL_PATH)


# Line Notify API URL
LINE_NOTIFY_API_URL = "https://notify-api.line.me/api/notify"

# Your Line token
LINE_TOKEN = "MHAk5pJRIm6daC3BLL8cT1guPQQxGNjIr1EgTOMqHgc"
line = LINE('MHAk5pJRIm6daC3BLL8cT1guPQQxGNjIr1EgTOMqHgc')
# Start video capture
cap = cv2.VideoCapture(-1)

while True:
    start_time = time.time()
    # Capture frame-by-frame
    ret, frame = cap.read()
    # line.sendimage(frame[:, :, ::-1])
   
    # time.sleep(5)

    if ret:
        # Use the FaceRecognition class to make predictions on the frame
        result = fr.predict(frame, threshold=0.5)

        if len(result['predictions']) > 0:
            # Get the name of the person recognized in the frame
            name = result['predictions'][0]['person']
            confidence = result['predictions'][0]['confidence']

            # Convert the frame to a base64-encoded string
            frame_b64 = base64.b64encode(frame).decode()

            # Send the results to Line
            headers = {
                "Authorization": "Bearer " + LINE_TOKEN,
                "Content-Type": "application/x-www-form-urlencoded"
            }
            payload = {
                "message": f"Person recognized: {name} ({confidence:.2f})",
                "imageFile": frame_b64
            }
            line.sendimage(frame[:, :, ::-1])
            response = requests.post(LINE_NOTIFY_API_URL, headers=headers, data=payload)
            pprint(response.json())
            
           
            current_time = time.time()
            if current_time - start_time >= 5:
                break
            print("Time's up")
    
    if cv2.waitKey(1) & 0xFF == ord('q'):
        break

# Release the video capture
cap.release()
cv2.destroyAllWindows()