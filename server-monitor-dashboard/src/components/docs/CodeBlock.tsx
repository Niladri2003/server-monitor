import React from 'react';
import { Light as SyntaxHighlighter } from 'react-syntax-highlighter';
import { docco } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import { Copy } from 'lucide-react';

interface CodeBlockProps {
  code: string;
  language: string;
}

export const CodeBlock: React.FC<CodeBlockProps> = ({ code, language }) => {
  const copyToClipboard = () => {
    navigator.clipboard.writeText(code);
  };

  return (
    <div className="relative group">
      <button
        onClick={copyToClipboard}
        className="absolute right-2 top-2 p-2 rounded-md bg-gray-800 text-gray-400 opacity-0 group-hover:opacity-100 transition-opacity"
      >
        <Copy className="w-4 h-4" />
      </button>
      <SyntaxHighlighter
        language={language}
        style={docco}
        className="rounded-lg !bg-gray-900 !p-4"
      >
        {code}
      </SyntaxHighlighter>
    </div>
  );
};