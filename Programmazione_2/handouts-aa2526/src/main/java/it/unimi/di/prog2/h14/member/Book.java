/*

Copyright 2025 Massimo Santini

This file is part of "Programmazione 2 @ UniMI" teaching material.

This is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This material is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this file.  If not, see <https://www.gnu.org/licenses/>.

*/

package it.unimi.di.prog2.h14.member;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Iterator;
import java.util.List;
import java.util.Objects;
import java.util.StringJoiner;

/**
 * A (promise of a) book, made up of chapters.
 *
 * <p>A book is the promise of an author to its editor, the author will first establish a title,
 * then he will add chapters to the book.
 */
public class Book implements Iterable<Book.Chapter> {

  /** The title of this book. */
  private final String title;

  /** The chapters of this book. */
  private final List<Chapter> chapters;

  /*
   * RI:
   *
   * - title is not null and not empty
   * - chapters is not null, does not contain null
   * - chapters contains only chapters whose enclosing book is this
   *
   * AF:
   *
   *  - title and chapters are the title and chapters of this book
   */

  /** A chapter of a book. */
  public class Chapter {

    /** The title of this chapter. */
    private final String title;

    /**
     * Creates a chapter with the given title.
     *
     * @param title the title of this chapter.
     * @throws NullPointerException if title is null.
     * @throws IllegalArgumentException if title is empty.
     */
    private Chapter(
        String title) { // private means it can be instantiated only by the enclosing Book
      if (Objects.requireNonNull(title).isEmpty()) throw new IllegalArgumentException();
      this.title = title;
    }

    /**
     * Returns the title of this chapter.
     *
     * @return the title.
     */
    public String title() {
      return title;
    }

    /**
     * Returns the book of which this chapter is a part.
     *
     * @return the book.
     */
    public Book book() {
      return Book.this;
    }

    @Override
    public String toString() {
      return Book.this.title + ": " + title;
    }
  }

  /**
   * Creates a book with the given title.
   *
   * <p>The book initially has no chapters.
   *
   * @param title the title of this book.
   * @throws NullPointerException if title is null.
   * @throws IllegalArgumentException if title is empty.
   */
  public Book(String title) {
    if (Objects.requireNonNull(title).isEmpty()) throw new IllegalArgumentException();
    this.title = title;
    this.chapters = new ArrayList<>();
  }

  /**
   * Adds a chapter with the given title to this book.
   *
   * @param title the title of the chapter to add.
   * @throws NullPointerException if title is null.
   * @throws IllegalArgumentException if title is empty.
   */
  public void addChapter(String title) {
    chapters.add(new Chapter(title));
  }

  @Override
  public Iterator<Chapter> iterator() {
    return Collections.unmodifiableCollection(chapters).iterator();
  }

  @Override
  public String toString() {
    StringJoiner sj = new StringJoiner("\n- ", title + "\n- ", "");
    for (Chapter chapter : chapters) {
      sj.add(chapter.title());
    }
    return sj.toString();
  }
}
