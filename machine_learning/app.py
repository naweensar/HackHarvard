from ultralytics import YOLO

# Load the trained model (from the saved checkpoint)
model = YOLO("trained_model.pt")


# Define a function to use the trained model for inference
def run_inference(input_path):
    # Call the model on the input (image or video path)
    results = model(input_path)

    # Display the results (bounding boxes, predictions, etc.)
    results[0].show()

    return results



# Example of calling the trained model for inference
video_path = "C:/Users/Naween/Pictures/Camera Roll/WIN_20241012_12_03_37_Pro.mp4"
results = run_inference(video_path)
