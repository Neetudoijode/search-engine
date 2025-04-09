import React from "react";

const ResultsList = ({ results, meta }) => {
  return (
    <div>
      <p className="mb-2">Found {meta.count} results in {meta.duration}</p>
      <ul className="space-y-2">
        {results.map((r, idx) => (
          <li key={idx} className="border p-2 rounded">
            <strong>{r.Message}</strong>
            <p>{r.Sender} - {r.EventId}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ResultsList;
