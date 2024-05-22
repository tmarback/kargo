import { ExtensionKind } from '@ui/features/extensions/extensions';

export const TestCreateStageExtension = {
  component: () => {
    return <div>Hello World</div>;
  },
  version: '0.0.1',
  kind: ExtensionKind.CreateStage,
  name: 'test',
  label: 'Test Extension'
};
