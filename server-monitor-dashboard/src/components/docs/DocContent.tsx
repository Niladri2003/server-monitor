import React from 'react';
import { useParams } from 'react-router-dom';
import { docContent } from './docContent';

export const DocContent = () => {
  const {  '*': subPath } = useParams();
  const path = subPath ? `${subPath}` : "docs";
  // console.log(path)
  // console.log(docContent)
  const content = docContent[path] || docContent['404'];

  return (
    <div className="prose prose-indigo max-w-none" dangerouslySetInnerHTML={{ __html: content }} />
  );
};