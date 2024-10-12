from ultralytics import YOLO

model = YOLO("trained_model2.pt")

results = model.track(source=0, show=True, tracker="bytetrack.yaml")