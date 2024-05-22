import {
  faChevronDown,
  faCog,
  faExternalLinkAlt,
  faPencil,
  faTrash
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Button, Dropdown } from 'antd';

import { useModal } from '@ui/features/common/modal/use-modal';
import { ExtensionKind, useExtensions } from '@ui/features/extensions/extensions';

import { DeleteProjectModal } from './components/delete-project-modal';
import { EditProjectModal } from './components/edit-project-modal';

export const ProjectSettings = () => {
  const { show: showEditModal } = useModal(EditProjectModal);
  const { show: showDeleteModal } = useModal(DeleteProjectModal);
  const extensions = useExtensions(ExtensionKind.ProjectAction);
  const { show: showExtension } = useModal();

  return (
    <Dropdown
      menu={{
        items: [
          {
            key: '1',
            label: (
              <>
                <FontAwesomeIcon icon={faPencil} size='xs' className='mr-2' /> Edit
              </>
            ),
            onClick: () => showEditModal()
          },
          {
            key: '2',
            danger: true,
            label: (
              <>
                <FontAwesomeIcon icon={faTrash} size='xs' className='mr-2' /> Delete
              </>
            ),
            onClick: () => showDeleteModal()
          },
          ...Object.values(extensions || {}).map((extension) => ({
            key: extension.name,
            label: (
              <>
                <FontAwesomeIcon
                  icon={extension.icon || faExternalLinkAlt}
                  size='xs'
                  className='mr-2'
                />{' '}
                {extension.label || extension.name}
              </>
            ),
            onClick: () => showExtension((props) => <extension.component {...props} />)
          }))
        ]
      }}
      placement='bottomRight'
      trigger={['click']}
    >
      <Button icon={<FontAwesomeIcon icon={faCog} size='1x' />}>
        <FontAwesomeIcon icon={faChevronDown} size='xs' />
      </Button>
    </Dropdown>
  );
};
