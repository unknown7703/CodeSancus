import React, { useState } from "react";
const UploadBox = () => {
  const [file, setFile] = useState();
  const handleFileSelect = (event) => {
    if (event.target.files[0]) {
      setFile(event.target.files[0]);
    }
  };
  const handleUpload = () => {
    if (!file) {
      return;
    }
      console.log("upload");
    
  };
  return (
    <div className="rounded-md text-black mt-40 ml-20 max-w-60 flex flex-col gap-3  ">
      <div className="bg-blue-800 h-1/6 p-2 rounded-md">
        <p className="text-white">Upload Folder</p>
      </div>
      <div className="h-5/6">
        <div>
          <input onChange={handleFileSelect} type="file" accept=".zip" />
        </div>
        <div>
          <button
            onClick={() => {
              handleUpload();
            }}
          >
            Upload
          </button>
        </div>
      </div>
    </div>
  );
};
export default UploadBox;
