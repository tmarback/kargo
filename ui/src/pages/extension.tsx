import { useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

import { ExtensionKind, useExtensions } from '@ui/features/extensions/extensions';

export const Extension = () => {
  const { name } = useParams();
  const extensions = useExtensions(ExtensionKind.SystemTab);
  const navigate = useNavigate();

  useEffect(() => {
    if (!name || !extensions || !extensions[name]) {
      navigate('/');
    }
  }, []);

  if (name && extensions && extensions[name]) {
    const Extension = extensions[name].component;
    return <Extension />;
  } else {
    return null;
  }
};
