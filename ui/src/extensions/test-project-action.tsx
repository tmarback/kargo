import { faDog } from '@fortawesome/free-solid-svg-icons';
import { Modal } from 'antd';

import { Extension, ExtensionKind, ProjectActionProps } from '@ui/features/extensions/extensions';

export const TestProjectActionExtension: Extension<ExtensionKind.ProjectAction> = {
  component: (props: ProjectActionProps) => (
    <Modal
      {...props}
      onCancel={props.hide}
      onOk={() => {
        alert('hello world!');
        props.hide();
      }}
    >
      Hello, World!
    </Modal>
  ),
  version: '0.0.1',
  kind: ExtensionKind.ProjectAction,
  name: 'TestProjectAction',
  label: 'My Cool Action',
  icon: faDog
};
