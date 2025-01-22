import React, { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';

interface TOCItem {
  id: string;
  text: string;
  level: number;
}

export const TableOfContents = () => {
  const [headings, setHeadings] = useState<TOCItem[]>([]);
  const [activeId, setActiveId] = useState('');
  const location = useLocation();

  useEffect(() => {
    const elements = Array.from(document.querySelectorAll('h1, h2, h3'))
      .map((element) => ({
        id: element.id,
        text: element.textContent || '',
        level: Number(element.tagName.charAt(1)),
      }));
    setHeadings(elements);

    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            setActiveId(entry.target.id);
          }
        });
      },
      { rootMargin: '0px 0px -80% 0px' }
    );

    elements.forEach(({ id }) => {
      const element = document.getElementById(id);
      if (element) {
        observer.observe(element);
      }
    });

    return () => observer.disconnect();
  }, [location]);

  return (
    <nav className="space-y-2">
      {headings.map((heading) => (
        <a
          key={heading.id}
          href={`#${heading.id}`}
          className={`block text-sm py-1 pl-${(heading.level - 1) * 4} ${
            activeId === heading.id
              ? 'text-indigo-600 font-medium'
              : 'text-gray-600 hover:text-gray-900'
          }`}
        >
          {heading.text}
        </a>
      ))}
    </nav>
  );
};