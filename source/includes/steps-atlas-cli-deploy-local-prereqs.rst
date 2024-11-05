
.. procedure:: 
   :style: normal 

   .. step:: Install the dependencies.

      a. Install the :ref:`{+atlas-cli+} <install-atlas-cli>`.

         If you use `Homebrew <https://brew.sh/#install>`__, you can
         run the following command in your terminal:

         .. code-block::

            brew install mongodb-atlas-cli
               
         For install instructions on other operating systems,
         see :ref:`install-atlas-cli`.

      #. Install `Docker <https://www.docker.com/>`__.

         - For MacOS or Windows, install `Docker Desktop v4.31+ <https://docs.docker.com/desktop/release-notes/#4310>`__. 
         - For Linux, install `Docker Engine v27.0+ <https://docs.docker.com/engine/release-notes/27.0/>`__.

         .. note::

            Docker requires a network connection for pulling and caching 
            MongoDB images.

            `Podman v5.0+ <https://podman.io>`__ is also supported for Linux RHEL versions.  

      #. (Optional) Install :mongosh:`mongosh </install>` version 2.0 or later.

         .. code-block:: sh

            brew install mongosh

         For install instructions on other operating 
         systems, see :mongosh:`Install mongosh </install>`.
         
      #. (Optional) Install :compass:`Compass </install>` version 1.39.4 or later.

         .. code-block:: sh

            brew install mongodb-compass

         For install instructions on other operating 
         systems, see :compass:`Download and Install Compass </install>`.
          
   .. step:: Authenticate to your |service| account.

      a. If you don't have an existing |service| account, run ``atlas setup`` in your terminal
         or `create a new account <https://account.mongodb.com/account/register?tck=docs_atlas>`__.

      #. In your terminal, run ``atlas auth login`` to authenticate with your 
         |service| login credentials. To learn more, see 
         :ref:`connect-atlas-cli`.
