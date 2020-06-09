import React, { useState, useMemo, useCallback } from 'react';
import Upvoter from "../Upvoter";
import styles from './index.module.scss';

const voteURL = "http://localhost:8001/api/v1/ranks/vote"

const Article: React.FC<{
  article: IArticle;
  user: IUser | undefined;
}> = ({
  article,
  user,
}) => {
    const articleVotes = useMemo(() => localStorage.getItem(article.ID.toString()) || "[]", [article]);
    const [votes, setVotes] = useState<Array<string>>(JSON.parse(articleVotes));
    const source = useMemo(() => {
      if(article.URL) {
        const url = new URL(article.URL);
        return url.hostname;
      }
    }, [
      article,
    ]);

    const {
      publishedDate,
      publishedTime,
    } = useMemo(() => {
      const publishedAt = new Date(article.PublishedAt)
      return {
        publishedDate: publishedAt.toLocaleDateString(),
        publishedTime: publishedAt.toLocaleTimeString(),
      }
    }, [
      article,
    ]);

    const canVote = useMemo<boolean | undefined>(() => (
      Boolean(user)
    ), [
      user
    ]);

    const hasVoted = useMemo<boolean | undefined>(() => (
      user && votes.includes(user.ID.toString())
    ), [
      votes,
      user,
    ]);

    const handleVote = useCallback(() => {
      if(user) {
        const options = {
          method: "POST",
          body: JSON.stringify({
            article: `${article.ID}`,
            user: `${user.ID}`,
          }),
        }

        fetch(voteURL, options)
          .then((res) => res.json())
          .then((users) => {
            localStorage.setItem(article.ID.toString(), JSON.stringify(users || []))
            setVotes(users || []);
          });
      }
    }, [
      article,
      user,
    ]);

    return (
      <div className={styles.Article}>
        <Upvoter
          className={styles.ArticleUpvoter}
          hasVoted={hasVoted}
          handleVote={handleVote}
          canVote={canVote}
        />
        <img className={styles.ArticleImage} src={article.Thumbnail} />
        <div className={styles.ArticleText}>
          <a className={styles.ArticleLink} href={article.URL} rel="noopener noreferrer" target="_blank">{article.Title}</a>
          <div className={styles.ArticleSmallText}>
            <a className={styles.ArticleSource} href={`https://${source}`} rel="noopener noreferrer" target="_blank">{source}</a>
            <span className={styles.ArticlePublishedDate}>{publishedDate}</span>
            <span className={styles.ArticlePublishedTime}>{publishedTime}</span>
          </div>
        </div>
      </div>
    );
  }

export default Article;
