from flask import Flask, jsonify
from ultralytics import YOLO
import threading

app = Flask(__name__)
model = YOLO("machine_learning/models/trained_model2.pt")

# Flag to control video tracking
tracking_active = False

@app.route('/start-video-feed', methods=['GET'])
def start_video_feed():
    global tracking_active
    if not tracking_active:
        tracking_thread = threading.Thread(target=track_video_feed)
        tracking_thread.start()
        tracking_active = True
        return jsonify({'success': True})
    return jsonify({'success': False, 'message': 'Video feed already running'})

def track_video_feed():
    global tracking_active
    try:
        model.track(source=0, show=True, tracker="bytetrack.yaml")
    except Exception as e:
        print(f"Error occurred: {e}")
    finally:
        tracking_active = False

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5000)