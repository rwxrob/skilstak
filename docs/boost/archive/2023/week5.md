# 2023 Boost - Week 5

📺 <https://youtu.be/ZcEIJFbHRn4>

This week we do "destructive" file system stuff related and take a lot of time to review the most common dangers. Watch this is you want to *keep* your job once you get it.

1. How do I create a new directory?
   
    ```sh
    mkdir NAME
    ```

1. How do I "hide" a file or directory?

    Give it a dot (`.`) as the first character in name.

1. How do I move or rename a file or directory?

    `mv OLD NEW`

    !!! danger
        Be very careful when moving and renaming since you can easily clobber an existing file or directory. Don't bother with activating `noclobber` because it will get in your way more than help you, just become aware/paranoid about the danger.

1. How do I move multiple files into a single directory?

    ```sh
    mv SOURCE... DIR
    ```

    SOURCE can be a file, directory (or path to) or a glob.

    !!! danger
        When moving muliple SOURCEs make *sure* that DIR is a directory and not a file.

    !!! note
        Note that `mv` sometimes will not work because the source and destination are on different file systems.

1. How do I list properties of a directory (instead of contents)?

    ```sh
    ls -lda DIR
    ```

1. How do I safely avoid typing long or complicated names?

    Tap tab (once or twice) for tab completion.

1. How do I remove a file?

    ```sh
    rm FILE
    ```

    !!! tip
        When working with destructive commands that operate on many files try `ls` first to make sure it only affects the files you want.

1. How do I create a temporary container (sandbox) just to try dangerous (stupid) things?

    ```sh
    podman run -it --rm ghcr.io/rwxrob/ws-skilstak
    ```

1. How do I remove a directory?

    ```sh
    rmdir DIR
    ```

    !!! note
        Prefer `rmdir` over alternatives to force yourself to be sure of each individual thing that is being deleted within that directory (since `rmdir` will not remove directories containing).

1. How do I remove a directory and everything under it recursively?

    ```sh
    rm -rf DIR
    ```

    !!! danger
        Even though this command does not work on significant system directories (like `/`) it is very recursively destructive and should almost never be used. Use `rmdir` instead.

1. How do I recover a broken or deleted configuration file?

    ```sh
    cp /etc/skel/FILE ~
    ```

1. How do I start a new shell and replace the currently running one?

    ```sh
    exec bash -l
    ```

1. How do I copy a file?

    ```sh
    cp FILE NEWNAME
    ```

1. How do I recusively copy a directory and subdirectories?

    ```sh
    cp -r SOURCE.. DIR
    ```

    !!! note
        This creates new files with new creation dates/times and ownership if copying from another user. Use other `cp` argument options when preserving them is important.
