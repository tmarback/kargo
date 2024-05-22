import { faAtom, faFlask } from '@fortawesome/free-solid-svg-icons';

import { ExtensionKind } from '@ui/features/extensions/extensions';

export const TestProjectTabExtension = {
  component: () => {
    return <div className='p-4'>Hello World</div>;
  },
  version: '0.0.1',
  kind: ExtensionKind.ProjectTab,
  name: 'test',
  label: 'Test Project Tab',
  icon: faFlask
};

export const AnotherTestProjectTabExtension = {
  component: () => {
    return <div className='p-4'>Another test extension</div>;
  },
  version: '0.0.1',
  kind: ExtensionKind.ProjectTab,
  name: 'another-test',
  label: 'Another Tab',
  icon: faAtom
};
