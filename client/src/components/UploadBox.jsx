import React, { useState } from "react";
import useApi from "../hooks/useApi";
import { API_URLS } from "../services/api.url";
const UploadBox = () => {
  const [file, setFile] = useState();
  const [name, setName] = useState();
  const uploadService = useApi(API_URLS.analyseDataSet);
  const handleNameChange = (event) => {
    setName(event.target.value);
  };
  const handleFileSelect = (event) => {
    if (event.target.files[0]) {
      setFile(event.target.files[0]);
    }
  };
  const handleUpload = () => {
    if (!file) {
      return;
    }
  
    const payload = {
      dataSetName: name,
      dataSetFile: file[0],
    };
    uploadService.call(payload);
    setName("");
    setFile("");
    document.getElementById("zipUp").value = "";
    document.getElementById("zipText").value = "";
  };

  return (
    <div className="rounded-md text-black  max-w-80 flex flex-col gap-3  ">
      <div className="bg-blue-800 h-1/6 p-2 rounded-md">
        <p className="text-white">Upload Folder</p>
      </div>
      <div className="h-5/6">
        <div>
          <input
            id="zipText"
            onChange={handleNameChange}
            type="text"
            placeholder="Give name to Dataset"
            className="border-2 border-black p-2 rounded-md mb-2 w-full"
          />
        </div>
        <div>
          <input
            id="zipUp"
            onChange={handleFileSelect}
            type="file"
            accept=".zip"
            className="border-2 border-black p-2 rounded-md mb-2 w-full"
          />
        </div>
        <div>
          <button
            className="mt-5 text-black bg-white hover:bg-blue-800  font-medium rounded-lg text-sm px-4 py-2 text-center dark:border-2 dark:border-black dark:hover:border-blue-800 dark:bg-white dark:hover:bg-blue-700 dark:hover:text-white "
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
