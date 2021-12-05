import querystring from "querystring";

const client_id = "140526a5c3074a7cba86129838ea667b";
const client_secret = "f819e6031f2b4ab58e6382212eede196";
const refresh_token =
  "AQAfSpxhdj5URPVooNoUN8RxDVPOKIcZLgkPNfwkJ7V8yCremIt05sW2OMgVec7l2Cft3uln6gvRNoNLsC9kQcg__fbr81GgmOhQ2i5OhjFYsKzk1eIwZ8Q8T85CRciII_4";

const basic = Buffer.from(`${client_id}:${client_secret}`).toString("base64");
const TOKEN_ENDPOINT = `https://accounts.spotify.com/api/token`;

const getAccessToken = async () => {
  const response = await fetch(TOKEN_ENDPOINT, {
    method: "POST",
    headers: {
      Authorization: `Basic ${basic}`,
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: querystring.stringify({
      grant_type: "refresh_token",
      refresh_token,
    }),
  });

  return response.json();
};

const TOP_TRACKS_ENDPOINT = `https://api.spotify.com/v1/me/top/tracks?time_range=long_term`;

export const getTopTracks = async () => {
  const { access_token } = await getAccessToken();

  return fetch(TOP_TRACKS_ENDPOINT, {
    headers: {
      Authorization: `Bearer ${access_token}`,
    },
  });
};

const CURRENTLY_PLAYING_ENDPOINT = `https://api.spotify.com/v1/me/player/currently-playing`;

export const getCurrentlyPlaying = async () => {
  const { access_token } = await getAccessToken();

  return fetch(CURRENTLY_PLAYING_ENDPOINT, {
    headers: {
      Authorization: `Bearer ${access_token}`,
    },
  });
};
