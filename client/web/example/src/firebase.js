// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import { getStorage } from 'firebase/storage';
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig =
{
  apiKey: "AIzaSyDAOgpsLn2edRfUBSf7pTfXhOk_EMbJ6Is",
  authDomain: "arche-6bd39.firebaseapp.com",
  projectId: "arche-6bd39",
  storageBucket: "arche-6bd39.appspot.com",
  messagingSenderId: "793359777992",
  appId: "1:793359777992:web:df040d93a1719df41265fd",
  measurementId: "G-JGMC21VFLM"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);
export const storage = getStorage(app);