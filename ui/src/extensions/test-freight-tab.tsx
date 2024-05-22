import { faHippo } from '@fortawesome/free-solid-svg-icons';

import { ExtensionKind } from '@ui/features/extensions/extensions';

export const TestFreightTabExtension = {
  name: 'test-freight-tab',
  component: () => {
    return <div>Test Freight Tab</div>;
  },
  label: 'Mud Bath',
  version: '0.0.1',
  icon: faHippo,
  kind: ExtensionKind.FreightTab
};
