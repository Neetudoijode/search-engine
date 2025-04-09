import React, { useState } from "react";
import SearchBox from "./components/SearchBox";
import ResultsList from "./components/ResultsList";

const App = () => {
  const [results, setResults] = useState([]);
  const [meta, setMeta] = useState({ count: 0, duration: "" });

  const handleSearch = async (keyword) => {
    const res = await fetch(`/search?q=${keyword}`);
    const data = await res.json();
    setResults(data.results);
    setMeta({ count: data.count, duration: data.duration });
  };

  return (
    <div className="p-4">
      <SearchBox onSearch={handleSearch} />
      <ResultsList results={results} meta={meta} />
    </div>
  );
};
export default App;