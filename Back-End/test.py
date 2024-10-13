import json
from inference_sdk import InferenceHTTPClient

# Initialize the inference client
CLIENT = InferenceHTTPClient(
    api_url="https://detect.roboflow.com/",
    api_key="pEWSO1IdAhzjOJSVVWr4"
)

# Perform inference on the image
try:
    result = CLIENT.infer("Output/upload-3521894298.jpg", model_id="human-fall-e2evv/2")

    # Check if the result is valid and contains predictions
    if result and isinstance(result, dict):
        predictions = result.get('predictions', [])
        print(predictions)
        # Log success and save the result to a JSON file
        print("Inference successful. Saving results...")
        print("confidence: ",result["predictions"][0]["confidence"])
        output_file = "inference_result.json"
        with open(output_file, "w") as f:
            json.dump(result, f, indent=4)

        print(f"Results saved to {output_file}")
    else:
        print("Inference failed or returned invalid result:", result)

except Exception as e:
    print(f"Error during inference: {e}")
