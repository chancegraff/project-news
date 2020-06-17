/// <reference path="index.d.ts" />

import React, { useEffect, useState, useCallback } from "react";
import validate from "validate.js";
import Article from "../Components/Article";
import Account from "../Components/Account";
import styles from "./index.module.scss";

const LOCAL_USER = "news-upvoter-user";
const LOCAL_HASH = "news-upvoter-hash";

const userConstraints = {
  ID: {
    presence: true,
  },
  Email: {
    presence: true,
  }
};

const getLocalUser = () => {
  const usr = localStorage.getItem(LOCAL_USER);
  if(usr) {
    const parsed = JSON.parse(usr);

    if(!validate(parsed, userConstraints)) {
      return parsed
    }
  }
};

const setLocalUser = (usr: IUser | undefined) => {
  if(usr) {
    if(!validate(usr, userConstraints)) {
      localStorage.setItem(LOCAL_USER, JSON.stringify(usr));
    }
  } else {
    localStorage.setItem(LOCAL_USER, "");
  }
};

const getLocalHash = () => {
  const hash = localStorage.getItem(LOCAL_HASH);
  if(hash) {
    const parsed = JSON.parse(hash);

    if(!validate.isEmpty(parsed) && validate.isString(parsed)) {
      return parsed;
    }
  }
};

const setLocalAuth = (hash: IAuth) => {
  if(hash && !validate.isEmpty(hash) && validate.isString(hash)) {
    localStorage.setItem(LOCAL_HASH, JSON.stringify(hash));
  }
};

const articlesURL = "http://localhost:8000/api/v1/articles";
const generateURL = "http://localhost:8000/api/v1/token/generate";

const identifiers: IIdentifiers = {
  software: navigator.platform,
  browser: navigator.userAgent,
  language: navigator.language,
  width: window.screen.width.toString(),
  height: window.screen.height.toString(),
  colors: window.screen.colorDepth.toString(),
  pixels: window.screen.pixelDepth.toString(),
};

const App = () => {
  const [articles, setArticles] = useState<Array<IArticle>>([]);
  const [auth, setStateAuth] = useState<IAuth>();
  const [user, setStateUser] = useState<IUser>();

  useEffect(() => {
    const fn = async () => {
      const usr = getLocalUser();

      if(usr) {
        setStateUser(usr);
      }

      const ath = getLocalHash() || {};

      if(ath.hash) {
        setStateAuth(ath);
      } else {
        const generateOptions = {
          method: "POST",
          body: JSON.stringify(identifiers),
        };

        const rsp = await fetch(generateURL, generateOptions);
        const js = await rsp.json();

        ath.hash = js;

        setStateAuth(ath);
        setLocalAuth(ath);
      }
    };

    fn();
  }, [
  ]);

  useEffect(() => {
    const fn = async () => {
      if(auth && auth.hash && !articles.length) {
        console.log('Effecting')
        const rsp = await fetch(articlesURL, {
          headers: {
            "X-Token-Auth": `Bearer ${auth.hash}`
          }
        });
        const js = await rsp.json();

        setArticles(js);
      }
    };

    fn();
  }, [
    auth,
    articles,
  ]);

  const setUser = useCallback((usr: IUser | undefined) => {
    setStateUser(usr);
    setLocalUser(usr);
  }, []);

  return (
    <div className={styles.App}>
      <div className={styles.ArticleList}>
        {articles.map((article) => (
          <Article key={article.ID} article={article} user={user} />
        ))}
      </div>
      <div className={styles.AccountBubble}>
        <Account user={user} setUser={setUser} auth={auth} />
      </div>
    </div>
  );
}

export default App;
