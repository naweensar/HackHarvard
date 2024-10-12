from ultralytics import YOLO

# Load the base model
model = YOLO("yolo11n.pt")

# Train the model
train_results = model.train(
    data="C:/Users/Naween/Downloads/Fallen Detect.v4i.yolov8/data.yaml",
    epochs=100,
    imgsz=640,
    device="cpu"
)

# Save the trained model (e.g., saves to 'runs/train/exp/weights/best.pt' by default)
model.save("trained_model2.pt")