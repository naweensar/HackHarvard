import * as tf from '@tensorflow/tfjs-node';
import * as cocoSsd from '@tensorflow-models/coco-ssd';
import NodeWebcam from 'node-webcam';

// Webcam Configuration
const webcamOpts = {
  width: 640,
  height: 480,
  delay: 0,
  saveShots: true,
  output: 'jpeg',
  device: false,
  callbackReturn: 'location',
  verbose: false,
};
const Webcam = NodeWebcam.create(webcamOpts);

const loadAndDetectObjects = async () => {
  const model = await cocoSsd.load();
  console.log('COCO-SSD model loaded.');

  Webcam.capture('temp', async (err, data) => {
    if (err) {
      console.error('Error capturing image from webcam:', err);
      return;
    }

    const image = fs.readFileSync('temp.jpg');
    const decodedImage = tf.node.decodeImage(image);
    const predictions = await model.detect(decodedImage);
    console.log('Predictions: ', predictions);
  });
};

// Start video feed
setInterval(loadAndDetectObjects, 1000);
