import DynamicPageContent from './DynamicPageContent';
import { MenuProvider } from '@/contexts/menuContext';

const DynamicPage = ({ params }: { params: { slug: string, subSlug: string } }) => {
  return (
    <MenuProvider>
      <DynamicPageContent params={params} />
    </MenuProvider>
  );
};
export default DynamicPage;