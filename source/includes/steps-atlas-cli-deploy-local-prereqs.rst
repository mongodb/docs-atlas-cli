
.. procedure:: 
   :style: normal 

   .. step:: Install the dependencies.

      a. Install the :ref:`{+atlas-cli+} <install-atlas-cli>`, 
         :dbtools:`MongoDB Database Tools </installation/installation/>`,
         and :mongosh:`mongosh </install>` version 2.0 or later.

         If you use `Homebrew <https://brew.sh/#install>`__, you can
         run the following commands in your terminal:

         .. code-block::

            brew install mongodb-atlas-cli
            brew tap mongodb/brew && brew install mongodb-database-tools
            brew install mongosh
         
         .. note:: 
               
            For install instructions on other operating systems,
            see the following resources:

            - :ref:`install-atlas-cli`
            - :dbtools:`Installing the Database Tools 
              </installation/installation/>`
            - :mongosh:`Install mongosh </install>`
            
      #. Install `Docker <https://www.docker.com/>`__.

         - For MacOS or Windows, install `Docker Desktop v4.31+ <https://docs.docker.com/desktop/release-notes/#4310>`__. 
         - For Linux, install `Docker Engine v27.0+ <https://docs.docker.com/engine/release-notes/27.0/>`__.

         .. note::

            Docker requires a network connection for pulling and caching 
            MongoDB images.

            `Podman v5.0+ <https://podman.io>`__ is also supported for Linux RHEL versions.  

      #. (Optional) Install :compass:`Compass </install>` version 1.39.4 or later.

         .. code-block:: sh

            brew install mongodb-compass

         For Compass install instructions on other operating 
         systems, see :compass:`Download and Install Compass </install>`.
          
   .. step:: Authenticate to your |service| account.

      a. If you don't have an existing |service| account, run ``atlas setup`` in your terminal
         or `create a new account <https://account.mongodb.com/account/register?tck=docs_atlas>`__.

      #. In your terminal, run ``atlas auth login`` to authenticate with your 
         |service| login credentials. To learn more, see 
         :ref:`connect-atlas-cli`.
