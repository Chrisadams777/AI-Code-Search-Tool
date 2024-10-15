// SearchBox.tsx - React component for search form
import React, { useState } from 'react';

type SearchBoxProps = {
  onSearch: (query: string, type: string) => void;
};

const SearchBox: React.FC<SearchBoxProps> = ({ onSearch }) => {
  const [query, setQuery] = useState('');
  const [type, setType] = useState('keyword');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSearch(query, type);
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        placeholder="Enter search query"
        value={query}
        onChange={(e) => setQuery(e.target.value)}
      />
      <select value={type} onChange={(e) => setType(e.target.value)}>
        <option value="keyword">Keyword</option>
        <option value="regex">Regex</option>
        <option value="semantic">Semantic</option>
      </select>
      <button type="submit">Search</button>
    </form>
  );
};

export default SearchBox;
