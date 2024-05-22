import { faCat } from '@fortawesome/free-solid-svg-icons';

import { ExtensionKind } from '@ui/features/extensions/extensions';

export const TestSystemTabExtension = {
  component: () => {
    return <div className='p-4'>Hello World</div>;
  },
  version: '0.0.1',
  kind: ExtensionKind.SystemTab,
  name: 'test',
  label: 'Demo',
  icon: faCat
};
