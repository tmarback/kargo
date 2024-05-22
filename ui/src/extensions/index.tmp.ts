/********
 * In this file, register all custom extensions by calling the registerExtension function.
 * Be sure to call registerExtension once for each extension _within_ the registerExtensions function.
 * For most extension kinds, you may register multiple extensions, as long as they have unique names.
 * Please note that after updating this file, you will need to rebuild the UI or restart the development server.
 *
 * All extensions must take the form:
 *
 * export const MyExtension = {
 *   component: (props: PairedType<T>) => React.ReactNode;
 *   version: string,
 *   kind: ExtensionKind,
 *   name: string,
 *   label? string,
 *   icon?: IconDefinition;
 * }
 *
 * Please refer to features/common/extensions.ts for available extension names and their corresponding props.
 * We recommend defining the extension in a separate file and importing it here.
 *****/

import { registerExtension } from '@ui/features/extensions/extensions';

import { TestCreateStageExtension } from './test-create-stage';
import { TestFreightTabExtension } from './test-freight-tab';
import { TestProjectActionExtension } from './test-project-action';
import { AnotherTestProjectTabExtension, TestProjectTabExtension } from './test-project-tab';
import { TestSystemTabExtension } from './test-system-tab';

export const registerExtensions = () => {
  // explicitly empty
  registerExtension(TestCreateStageExtension);
  registerExtension(TestProjectTabExtension);
  registerExtension(AnotherTestProjectTabExtension);
  registerExtension(TestSystemTabExtension);
  registerExtension(TestProjectActionExtension);
  registerExtension(TestFreightTabExtension);
};
