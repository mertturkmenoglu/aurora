import { createClient } from 'contentful';

export const CF_SPACE_ID = '32jmgckgm0ny';
export const CF_CDA_ACCESS_TOKEN = 'x_wBsN3NPk7KPeEZC3xUN-el_FlupndHfz4X_1ie7pY';

export const client = createClient({
  // This is the space ID. A space is like a project folder in Contentful terms
  space: CF_SPACE_ID,
  // This is the access token for this space. Normally you get both ID and the token in the Contentful web app
  accessToken: CF_CDA_ACCESS_TOKEN,
});

export async function getHomeBanners() {
  return client.getEntries({
    content_type: 'banner',
  });
}
