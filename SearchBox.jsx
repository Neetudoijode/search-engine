import React, { useState } from "react";

const SearchBox = ({ onSearch }) => {
  const [input, setInput] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    onSearch(input);
  };

  return (
    <form onSubmit={handleSubmit} className="mb-4">
      <input
        className="border p-2 mr-2"
        type="text"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="Enter search keyword..."
      />
      <button className="bg-blue-500 text-white px-4 py-2" type="submit">
        Search
      </button>
    </form>
  );
};

export default SearchBox;
