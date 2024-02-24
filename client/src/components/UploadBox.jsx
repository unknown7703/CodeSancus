import React, { useState } from "react";
const UploadBox = () => {
  const [file, setFile] = useState();
  const handleFileSelect = (event) => {
    if (event.target.file) {
      setFile(event.target.file[0]);
    }
  };
  const handleUpload = () => {
    if (!file) {
      return;
    }
    console.log("upload logic");
  };
  return (
    <div className="rounded-md drop-shadow-lg ">
      <div className="bg-blue-800 h-1/6">
        <p className="text-white">Upload Folder</p>
      </div>
      <div className="h-5/6">
        <div>
          <input onChange={handleFileSelect} type="file" accept=".zip" />
        </div>
        <div>
          <button onClick={handleUpload}>Uplaod</button>
        </div>
      </div>
    </div>
  );
};
export default UploadBox;
